g++ -fPIC -o glop_sdl.o -c -Iinclude `pkg-config --cflags sdl` glop_sdl.cpp
g++ -fPIC -shared -o libglop.so `pkg-config --libs sdl` glop_sdl.o
rm -f glop_sdl.o
mkdir -p lib
mv libglop.so lib
