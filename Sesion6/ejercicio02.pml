#define wait(s) atomic{s>0 -> s --}
#define signal(s) s++

#define N 10

byte buffer[N]
byte notEmpty = 0
byte notFull = N
byte limite = 0
byte mutex = 1

active proctype productor() {
    byte d
    do
    ::
        d ++
        wait(notFull)

        wait(mutex)
        buffer[limite] = d
        limite ++
        printf("Depositando mensaje en la cola!\n")
        signal(mutex)

        signal(notEmpty)
    od
}

active proctype consumidor() {
    byte d
    do
    ::
        wait(notEmpty)

        wait(mutex)
        limite --
        d = buffer[limite]
        signal(mutex)

        signal(notFull)
        printf("Recuperando mensaje de la cola!!!\n")
    od
}

init {
    run productor()
    run consumidor()
}