g++ -fPIC -o glop_x11.o -c -Iinclude glop_x11.cpp
g++ -fPIC -shared -o libglop.so glop_x11.o
rm -f glop_x11.o
mkdir -p lib
mv libglop.so lib
