// Lógica del canvas del grafo (equivalente al script inline de la versión Java).
// initGrafo(nodos, aristas) se llama desde index.html con los datos que
// el servidor Go inyectó en la plantilla.

function initGrafo(nodos, aristas) {
    const STORAGE_KEY = 'dijkstra-node-positions';
    const posicionesPorDefecto = {
        A: { x: 140, y: 90 },
        B: { x: 380, y: 180 },
        C: { x: 250, y: 250 },
        D: { x: 520, y: 250 },
        E: { x: 380, y: 360 },
        F: { x: 620, y: 360 }
    };
    const radioNodo = 24;

    let nodoArrastrado = null;
    let menuNodoActual = null;
    let posiciones = cargarPosiciones();

    function cargarPosiciones() {
        try {
            return JSON.parse(localStorage.getItem(STORAGE_KEY) || '{}');
        } catch (error) {
            return {};
        }
    }

    function guardarPosiciones() {
        localStorage.setItem(STORAGE_KEY, JSON.stringify(posiciones));
    }

    function obtenerPosicion(nombre) {
        if (!posiciones[nombre]) {
            const base = posicionesPorDefecto[nombre] || (() => {
                const indice = nodos.indexOf(nombre);
                const columnas = 3;
                const fila = Math.floor(indice / columnas);
                const columna = indice % columnas;
                return {
                    x: 140 + columna * 180,
                    y: 90 + fila * 140
                };
            })();
            posiciones[nombre] = { ...base };
            guardarPosiciones();
        }
        return posiciones[nombre];
    }

    function dibujarGrafo() {
        const canvas = document.getElementById('canvas');
        if (!canvas) return;
        const ctx = canvas.getContext('2d');
        const rect = canvas.getBoundingClientRect();
        canvas.width = rect.width;
        canvas.height = rect.height;

        ctx.clearRect(0, 0, canvas.width, canvas.height);

        ctx.strokeStyle = '#457b9d';
        ctx.lineWidth = 2;
        aristas.forEach(arista => {
            const origen = obtenerPosicion(arista.origen);
            const destino = obtenerPosicion(arista.destino);
            ctx.beginPath();
            ctx.moveTo(origen.x, origen.y);
            ctx.lineTo(destino.x, destino.y);
            ctx.stroke();

            const mx = (origen.x + destino.x) / 2;
            const my = (origen.y + destino.y) / 2;
            ctx.fillStyle = '#1d3557';
            ctx.font = '12px Arial';
            ctx.fillText(arista.peso, mx + 8, my - 8);
        });

        nodos.forEach(nombre => {
            const pos = obtenerPosicion(nombre);
            ctx.beginPath();
            ctx.arc(pos.x, pos.y, radioNodo, 0, Math.PI * 2);
            ctx.fillStyle = '#1d3557';
            ctx.fill();
            ctx.strokeStyle = '#ffffff';
            ctx.lineWidth = 2;
            ctx.stroke();
            ctx.fillStyle = '#ffffff';
            ctx.font = 'bold 14px Arial';
            ctx.textAlign = 'center';
            ctx.textBaseline = 'middle';
            ctx.fillText(nombre, pos.x, pos.y);
        });
    }

    function obtenerNodoEnCoordenadas(x, y) {
        for (let i = nodos.length - 1; i >= 0; i--) {
            const nombre = nodos[i];
            const pos = obtenerPosicion(nombre);
            const distancia = Math.hypot(x - pos.x, y - pos.y);
            if (distancia <= radioNodo + 4) {
                return nombre;
            }
        }
        return null;
    }

    function mostrarMenuNodo(nombre, x, y) {
        const menu = document.getElementById('nodeMenu');
        menuNodoActual = nombre;
        menu.style.display = 'block';
        menu.style.left = `${x + 10}px`;
        menu.style.top = `${y + 10}px`;
    }

    function ocultarMenuNodo() {
        const menu = document.getElementById('nodeMenu');
        if (menu) menu.style.display = 'none';
        menuNodoActual = null;
    }

    function manejarAccionMenu(action) {
        if (!menuNodoActual) return;
        if (action === 'delete') {
            const form = document.createElement('form');
            form.method = 'post';
            form.action = '/eliminar-nodo';
            const input = document.createElement('input');
            input.type = 'hidden';
            input.name = 'nombre';
            input.value = menuNodoActual;
            form.appendChild(input);
            document.body.appendChild(form);
            form.submit();
        }
        ocultarMenuNodo();
    }

    document.querySelectorAll('#nodeMenu button[data-action]').forEach(boton => {
        boton.addEventListener('click', () => manejarAccionMenu(boton.getAttribute('data-action')));
    });

    function manejarMouseDown(evento) {
        const rect = evento.currentTarget.getBoundingClientRect();
        const x = evento.clientX - rect.left;
        const y = evento.clientY - rect.top;
        const nodo = obtenerNodoEnCoordenadas(x, y);
        if (nodo) {
            nodoArrastrado = nodo;
            evento.currentTarget.style.cursor = 'grabbing';
        } else {
            ocultarMenuNodo();
        }
    }

    function manejarContextMenu(evento) {
        const rect = evento.currentTarget.getBoundingClientRect();
        const x = evento.clientX - rect.left;
        const y = evento.clientY - rect.top;
        const nodo = obtenerNodoEnCoordenadas(x, y);

        if (nodo) {
            evento.preventDefault();
            mostrarMenuNodo(nodo, x, y);
        } else {
            ocultarMenuNodo();
        }
    }

    function manejarMouseMove(evento) {
        if (!nodoArrastrado) return;
        const rect = evento.currentTarget.getBoundingClientRect();
        const x = evento.clientX - rect.left;
        const y = evento.clientY - rect.top;
        const pos = obtenerPosicion(nodoArrastrado);
        pos.x = x;
        pos.y = y;
        dibujarGrafo();
    }

    function manejarMouseUp() {
        if (nodoArrastrado) {
            guardarPosiciones();
            nodoArrastrado = null;
            const canvas = document.getElementById('canvas');
            if (canvas) canvas.style.cursor = 'grab';
        }
    }

    const canvas = document.getElementById('canvas');
    if (canvas) {
        canvas.addEventListener('mousedown', manejarMouseDown);
        canvas.addEventListener('mousemove', manejarMouseMove);
        canvas.addEventListener('mouseup', manejarMouseUp);
        canvas.addEventListener('mouseleave', manejarMouseUp);
        canvas.addEventListener('contextmenu', manejarContextMenu);
    }

    document.addEventListener('mousedown', (evento) => {
        const dentroMenu = evento.target.closest('#nodeMenu');
        const dentroGrafo = evento.target.closest('.grafo');
        if (!dentroMenu && !dentroGrafo) {
            ocultarMenuNodo();
        }
    });

    window.addEventListener('resize', dibujarGrafo);
    dibujarGrafo();
}