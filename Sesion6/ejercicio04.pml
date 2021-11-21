#define wait(s) atomic{s>0 -> s --}
#define signal(s) s++

#define N 5

byte fork[N] = {1,1,1,1,1}
byte room = 4

active[N] proctype filosofo() {
    pid id = _pid
    do
    ::
        wait(room)
        printf("Pensando!!!\n")
        wait(fork[id])
        wait(fork[(id + 1)%N])
        printf("Comiendo!!!\n")
        signal(fork[id])
        signal(fork[(id + 1)%N])
        signal(room)
    od
}