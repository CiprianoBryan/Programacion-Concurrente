byte turno=1

active proctype p() {
    do
    ::
        printf("SNC line1\n")
        printf("SNC line2\n")

        (turno == 1) ->
            printf("SC line1\n")
            printf("SC line2\n")
            turno = 2
    od
}

active proctype q() {
    do
    ::
        printf("SNC line1\n")
        printf("SNC line2\n")

        (turno == 2) ->
            printf("SC line1\n")
            printf("SC line2\n")
            turno = 1
    od
}