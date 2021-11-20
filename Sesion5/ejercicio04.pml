byte wantp = 0
byte wantq = 0
byte c = 0

active proctype p() {
    do
    ::
        (wantq == 0) ->
            wantp = 1
            c ++
            assert(c < 2)
            c --
            wantp = 0
    od
}

active proctype q() {
    do
    ::
        (wantp == 0) ->
            wantq = 1
            c ++
            assert(c < 2)
            c --
            wantq = 0
    od
}