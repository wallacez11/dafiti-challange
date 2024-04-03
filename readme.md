### 1 - Supone que en un repositorio GIT hiciste un commit y te olvidaste un archivo. Explicar como se soluciona si hiciste push, y como si aun no hiciste. De ser posible, buscar que quede solo un commit con los cambios.

```
Si lo has hecho push:
git add archivo-que-me-olvide.txt
git commit --amend --no-edit
git push --force-with-lease ==> es una opción más segura que no sobrescribirá ningún trabajo en la rama remota si se agregaron más confirmaciones a la rama remota (por otro miembro del equipo o compañero de trabajo o lo que sea). Garantiza que no sobrescriba el trabajo de otra persona al forzarlo
```

Si no has hecho push:

```
git add archivo-que-me-olvide.txt
git commit --amend --no-edit
```

### 2 - Tenes un sitio en tu computadora de desarrollo, y cuando entras desde el navegador, en la consola te aparece esto: “https://site.local/favicon.ico Failed to load resource: net::ERR_INSECURE_RESPONSE”. El archivo existe, el sitio debe cargar por https. Como lo solucionas?

Si el archivo existe, este error suele estar relacionado con la configuración del certificado SSL, como solución intentar corregir el certificado SSL del sitio web, comprobar también si el servidor está configurado para exportar recursos vía https y, por último, intentar comprobar si Se requiere cualquier configuración de Cors.

### 3 - Tenes un archivo comiteado en un repositorio GIT que deseas que quede ignorado. Que pasos debes realizar?

```
Agregaría el archivo a .gitignore
git rm --cached archivo.txt para eliminar el archivo del índice git
git commit -m "Dejar de rastrear el archivo archivo.txt"
git push

```

### 4 - Explica las ventajas de cargar en tu sitio las librerias de terceros por GTM.

Tag manager ofrece una plataforma que facilita la implementación y gestión de etiquetas de terceros utilizadas en el sitio web, dando un mayor control sobre los datos de acceso o conversión.

### 5 - Ejercicio

Archivo main.go

### 6 - Mysql

```
-- Consulta 1: ¿Cuál es el jugador más viejo de cada equipo?
SELECT e.nombre AS equipo, j.nombre AS jugador, j.fecha_nacimiento AS fecha_nacimiento
FROM equipos e
INNER JOIN jugadores j ON e.id_equipos = j.fk_equipos
WHERE (j.fecha_nacimiento, e.id_equipos) IN
    (SELECT MIN(j.fecha_nacimiento), e.id_equipos
    FROM jugadores j
    WHERE j.fk_equipos = e.id_equipos
    GROUP BY e.id_equipos);

-- Consulta 2: ¿Cuántos partidos jugó de visitante cada equipo?
SELECT e.nombre AS equipo, COUNT(p.id_partidos) AS partidos_de_visitante
FROM equipos e
LEFT JOIN partidos p ON e.id_equipos = p.fk_equipo_visitante
GROUP BY e.nombre;

-- Consulta 3: ¿Qué equipos jugaron el 01/01/2016 y el 12/02/2016?
SELECT e.nombre AS equipo
FROM equipos e
JOIN partidos p ON e.id_equipos = p.fk_equipo_local OR e.id_equipos = p.fk_equipo_visitante
WHERE p.fecha_partido BETWEEN '2016-01-01' AND '2016-12-02'
GROUP BY e.nombre;

-- Consulta 4: Diga el total de goles que hizo el equipo “Chacarita” en su historia (como local o visitante)
SELECT SUM(goles) AS total_goles
FROM (
    SELECT SUM(goles_local) AS goles
    FROM partidos
    WHERE fk_equipo_local = (SELECT id_equipos FROM equipos WHERE nombre = 'Equipo A')
    UNION ALL
    SELECT SUM(goles_visitante) AS goles
    FROM partidos
    WHERE fk_equipo_visitante = (SELECT id_equipos FROM equipos WHERE nombre = 'Equipo A')
)
```

### Extra

Una democratización más integral de los servicios de desarrollo con plataformas low code.
