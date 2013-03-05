g++ -fPIC -o glop_sdl.o -c -Iinclude `pkg-config --cflags sdl` glop_sdl.cpp
g++ -fPIC -shared -o libglop.so glop_sdl.o `pkg-config --libs sdl`
rm -f glop_sdl.o
mkdir -p lib
mv libglop.so lib
