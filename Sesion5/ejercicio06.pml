#define wait(s) atomic{s>0 -> s--}
#define signal(s) s++

byte S = 1

active[2] proctype p() {
    do
    ::
        printf("SNC line1\n")
        printf("SNC line2\n")
        wait(S)
        printf("SC line1\n")
        printf("SC line2\n")
        signal(S)
    od
}