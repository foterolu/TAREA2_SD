# Distribucion  y Ejecución de máquinas
- Maquina 1 : Nodo1
- Maquina 2: Nodo2
- Maquina 3: Nodo3
- Maquina 4: NameNode  y Cliente

al momento de entrar a cada maquina ejecutar Make run y adicionalmente, en la maquina 4 , abrir otra consola y ejecutar
- go run TAREA2_SD/cliente/cliente.go -cliente=uploader  (para subir el primer libro del directorio downloader)
- go run TAREA2_SD/cliente/cliente.go -cliente=downloader -libro=Diez_negritos-Christie_Agatha (que es el libro por defecto)


# Concurrencia
 Al momento de escribir o leer sobre un archivo, la funcion UploadChunk, que es la que invoca todas las otras funciones de distrbucion y generacion de propuesta, está Lockead por un RWsync.RWMutex.Lock() -> Unlock(), De manera que nunca dos GOroutines chocan al momento de acceder a estas funcionalidades.

 # Disclaimer
 El Proyecto funcionaba correctamente desde windows, es decir, los funcionalidades de uploader, downloader, distribuido y centralizado corren de manera adecuada, pero debido a la naturaleza de los makefiles, la funcion Iowrite (que usamos para crear los archivos binarios) no crea los archivos en las carpetas que deseamos, por lo que la funcionalidad de cliente downloader no es capaz de recopilar los archivos correctamente. Las pruebas de las cantidad de mensajes las probamos desde windows