byte wantp = 0
byte wantq = 0
byte c = 0

active proctype p() {
    do
    ::
        wantp = 1
        (wantq == 0) ->
            c ++
            assert(c < 2)
            c --
            wantp = 0
    od
}

active proctype q() {
    do
    ::
        wantq = 1
        (wantp == 0) ->
            c ++
            assert(c < 2)
            c --
            wantq = 0
    od
}