# test repo for the issue akiyosi/goneovim#434

# Step to reproduce the problem

## checkout neovim/neovim commit 1d16bba

    git checkout 1d16bba # Commits older than this will not reproduce 

And build neovim # see neovim build procedures.

## build test-go-nvim-ui

    go build main.go

## Reproduces the problem

When the program is executed as follows, 

    ./main &

the program does not exit and nvim is in `<defunct>` state, donfirming the xsel process is up and running.

    akiyosi@pop-os:~/test-go-nvim-ui$ ps -ef | grep nvim
    akiyosi   219795  219789  0 00:56 pts/1    00:00:00 [nvim] <defunct>
    akiyosi@pop-os:~/test-go-nvim-ui$ ps -ef | grep main
    akiyosi   219789  196252  0 00:56 pts/1    00:00:00 ./main
    akiyosi@pop-os:~/test-go-nvim-ui$ ps -ef | grep xsel
    akiyosi   219797    2003  0 00:56 ?        00:00:00 /usr/bin/xsel --nodetach -i -p


