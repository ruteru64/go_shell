module mygoshell.com/s

go 1.16

replace local.packages/shellerr => ./shellerr

replace local.packages/shell => ./shell

replace local.packages/a => ./a

replace local.packages/pwd => ./pwd

require golang.org/x/tools v0.1.8 // indirect
