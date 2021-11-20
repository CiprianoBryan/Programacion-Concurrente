byte turno = 1
byte c = 0

active proctype p() {
    do
    ::
        (turno == 1) ->
            c ++
            assert(c < 2)
            c --
            turno = 2
    od
}

active proctype q() {
    do
    ::
        (turno == 2) ->
            c ++
            assert(c < 2)
            c --
            turno = 1
    od
}