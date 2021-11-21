#define wait(s) atomic{s>0 -> s --}
#define signal(s) s++

#define N 5

byte fork[N] = {1,1,1,1,1}

active[N-1] proctype filosofo() {
    pid id = _pid
    do
    ::
        printf("Pensando!!!\n")
        wait(fork[id])
        wait(fork[(id + 1)%N])
        printf("Comiendo!!!\n")
        signal(fork[id])
        signal(fork[(id + 1)%N])
    od
}

active proctype filosofo_four() {
    do
    ::
        printf("Pensando!!!\n")
        wait(fork[0])
        wait(fork[4])
        printf("Comiendo!!!\n")
        signal(fork[0])
        signal(fork[4])
    od
}