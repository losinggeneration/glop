mkdir -p ../lib
g++ -fPIC -o glop_sdl.o `pkg-config --cflags sdl` -c glop_sdl.cpp
g++ -fPIC -shared -o ../lib/libglop.so `pkg-config --libs sdl` glop_sdl.o
rm -f glop_sdl.o
