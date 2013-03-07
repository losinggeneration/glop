mkdir -p ../lib
g++ -fPIC -o glop_x11.o -c glop_x11.cpp
g++ -fPIC -shared -o ../lib/libglop.so glop_x11.o -lX11 -lGL
rm -f glop_x11.o
