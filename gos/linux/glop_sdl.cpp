#include <vector>
#include <string>

#include "SDL.h"

using namespace std;

extern "C" {

#include "glop.h"

typedef short GlopKey;

void GlopInit() {
	SDL_Init(SDL_INIT_VIDEO);
}

void glopShutDown() {
	SDL_Quit();
}

vector<GlopKeyEvent> events;
static bool SynthKey(const SDLKey &sym, const SDL_Event &event, bool pushed, GlopKeyEvent *ev) {
  SDL_Event e;

  GlopKey ki = 0;
  switch(sym) {
    case SDLK_a: ki = SDLK_a; break;
    case SDLK_b: ki = SDLK_b; break;
    case SDLK_c: ki = SDLK_c; break;
    case SDLK_d: ki = SDLK_d; break;
    case SDLK_e: ki = SDLK_e; break;
    case SDLK_f: ki = SDLK_f; break;
    case SDLK_g: ki = SDLK_g; break;
    case SDLK_h: ki = SDLK_h; break;
    case SDLK_i: ki = SDLK_i; break;
    case SDLK_j: ki = SDLK_j; break;
    case SDLK_k: ki = SDLK_k; break;
    case SDLK_l: ki = SDLK_l; break;
    case SDLK_m: ki = SDLK_m; break;
    case SDLK_n: ki = SDLK_n; break;
    case SDLK_o: ki = SDLK_o; break;
    case SDLK_p: ki = SDLK_p; break;
    case SDLK_q: ki = SDLK_q; break;
    case SDLK_r: ki = SDLK_r; break;
    case SDLK_s: ki = SDLK_s; break;
    case SDLK_t: ki = SDLK_t; break;
    case SDLK_u: ki = SDLK_u; break;
    case SDLK_v: ki = SDLK_v; break;
    case SDLK_w: ki = SDLK_w; break;
    case SDLK_x: ki = SDLK_x; break;
    case SDLK_y: ki = SDLK_y; break;
    case SDLK_z: ki = SDLK_z; break;

    case SDLK_0: ki = SDLK_0; break;
    case SDLK_1: ki = SDLK_1; break;
    case SDLK_2: ki = SDLK_2; break;
    case SDLK_3: ki = SDLK_3; break;
    case SDLK_4: ki = SDLK_4; break;
    case SDLK_5: ki = SDLK_5; break;
    case SDLK_6: ki = SDLK_6; break;
    case SDLK_7: ki = SDLK_7; break;
    case SDLK_8: ki = SDLK_8; break;
    case SDLK_9: ki = SDLK_9; break;

    case SDLK_F1: ki = kKeyF1; break;
    case SDLK_F2: ki = kKeyF2; break;
    case SDLK_F3: ki = kKeyF3; break;
    case SDLK_F4: ki = kKeyF4; break;
    case SDLK_F5: ki = kKeyF5; break;
    case SDLK_F6: ki = kKeyF6; break;
    case SDLK_F7: ki = kKeyF7; break;
    case SDLK_F8: ki = kKeyF8; break;
    case SDLK_F9: ki = kKeyF9; break;
    case SDLK_F10: ki = kKeyF10; break;
    case SDLK_F11: ki = kKeyF11; break;
    case SDLK_F12: ki = kKeyF12; break;

    case SDLK_KP0: ki = kKeyPad0; break;
    case SDLK_KP1: ki = kKeyPad1; break;
    case SDLK_KP2: ki = kKeyPad2; break;
    case SDLK_KP3: ki = kKeyPad3; break;
    case SDLK_KP4: ki = kKeyPad4; break;
    case SDLK_KP5: ki = kKeyPad5; break;
    case SDLK_KP6: ki = kKeyPad6; break;
    case SDLK_KP7: ki = kKeyPad7; break;
    case SDLK_KP8: ki = kKeyPad8; break;
    case SDLK_KP9: ki = kKeyPad9; break;

    case SDLK_LEFT: ki = kKeyLeft; break;
    case SDLK_RIGHT: ki = kKeyRight; break;
    case SDLK_UP: ki = kKeyUp; break;
    case SDLK_DOWN: ki = kKeyDown; break;

    case SDLK_BACKSPACE: ki = kKeyBackspace; break;
    case SDLK_TAB: ki = kKeyTab; break;
    case SDLK_KP_ENTER: ki = kKeyPadEnter; break;
    case SDLK_RETURN: ki = kKeyReturn; break;
    case SDLK_ESCAPE: ki = kKeyEscape; break;

    case SDLK_LSHIFT: ki = kKeyLeftShift; break;
    case SDLK_RSHIFT: ki = kKeyRightShift; break;
    case SDLK_LCTRL: ki = kKeyLeftControl; break;
    case SDLK_RCTRL: ki = kKeyRightControl; break;
    case SDLK_LALT: ki = kKeyLeftAlt; break;
    case SDLK_RALT: ki = kKeyRightAlt; break;
    case SDLK_LSUPER: ki = kKeyLeftGui; break;
    case SDLK_RSUPER: ki = kKeyRightGui; break;

    case SDLK_KP_DIVIDE: ki = kKeyPadDivide; break;
    case SDLK_KP_MULTIPLY: ki = kKeyPadMultiply; break;
    case SDLK_KP_MINUS: ki = kKeyPadSubtract; break;
    case SDLK_KP_PLUS: ki = kKeyPadAdd; break;

    case SDLK_BACKQUOTE: ki = '`'; break;
    case SDLK_MINUS: ki = '-'; break;
    case SDLK_EQUALS: ki = '='; break;
    case SDLK_LEFTBRACKET: ki = '['; break;
    case SDLK_RIGHTBRACKET: ki = ']'; break;
    case SDLK_BACKSLASH: ki = '\\'; break;
    case SDLK_SEMICOLON: ki = ';'; break;
    case SDLK_QUOTE: ki = '\''; break;
    case SDLK_COMMA: ki = ','; break;
    case SDLK_PERIOD: ki = '.'; break;
    case SDLK_SLASH: ki = '/'; break;
    case SDLK_SPACE: ki = '/'; break;
  }

  if(ki == 0)
    return false;

  int x, y;

  SDL_GetMouseState(&x, &y);

  ev->index = ki;
  ev->press_amt = pushed ? 1.0 : 0.0;
  ev->timestamp = SDL_GetTicks();
  ev->cursor_x = x;
  ev->cursor_y = y;
  ev->num_lock = event.key.keysym.mod & KMOD_NUM;
  ev->caps_lock = event.key.keysym.mod & KMOD_CAPS;
  return true;
}
static bool SynthButton(int button, bool pushed, const SDL_Event &event, GlopKeyEvent *ev) {
  int x, y;

  SDL_GetMouseState(&x, &y);

  GlopKey ki;
  if(button == SDL_BUTTON(1))
    ki = kMouseLButton;
  else if(button == SDL_BUTTON(2))
    ki = kMouseMButton;
  else if(button == SDL_BUTTON(3))
    ki = kMouseRButton;
  /*
  else if(button == Button4)
    ki = kMouseWheelUp;   // these might be inverted so they're disabled for now
  else if(button == Button5)
    ki = kMouseWheelDown;*/
  else
    return false;

  ev->index = ki;
  ev->press_amt = pushed ? 1.0 : 0.0;
  ev->timestamp = SDL_GetTicks();
  ev->cursor_x = x;
  ev->cursor_y = y;
  ev->num_lock = event.key.keysym.mod & KMOD_NUM;
  ev->caps_lock = event.key.keysym.mod & KMOD_CAPS;
  return true;
}

static bool SynthMotion(int dx, int dy, const SDL_Event &event, GlopKeyEvent *ev, GlopKeyEvent *ev2) {
  int x, y;

  SDL_GetMouseState(&x, &y);

  ev->index = kMouseXAxis;
  ev->press_amt = dx;
  ev->timestamp = SDL_GetTicks();
  ev->cursor_x = x;
  ev->cursor_y = y;
  ev->num_lock = event.key.keysym.mod & KMOD_NUM;
  ev->caps_lock = event.key.keysym.mod & KMOD_CAPS;

  ev2->index = kMouseYAxis;
  ev2->press_amt = dy;
  ev2->timestamp = ev->timestamp;
  ev2->cursor_x = x;
  ev2->cursor_y = y;
  ev2->num_lock = event.key.keysym.mod & KMOD_NUM;
  ev2->caps_lock = event.key.keysym.mod & KMOD_CAPS;

  return true;
}

void GlopThink() {
	if(SDL_GetVideoSurface() == NULL) return;

	SDL_Event event;
	while(SDL_PollEvent(&event) != 0) {
		GlopKeyEvent ev;
		GlopClearKeyEvent(&ev);
		switch(event.type) {
			case SDL_KEYDOWN:
				if(SynthKey(event.key.keysym.sym, event, true, &ev))
					events.push_back(ev);
				break;
			case SDL_KEYUP:
				if(SynthKey(event.key.keysym.sym, event, false, &ev))
					events.push_back(ev);
				break;
			case SDL_MOUSEMOTION:
				GlopKeyEvent ev2;
				GlopClearKeyEvent(&ev2);
				if(SynthMotion(event.motion.xrel, event.motion.yrel, event, &ev, &ev2)) {
					events.push_back(ev);
					events.push_back(ev2);
				}
				break;
			case SDL_MOUSEBUTTONDOWN:
				if(SynthButton(event.button.button, true, event, &ev))
					events.push_back(ev);
				break;
			case SDL_MOUSEBUTTONUP:
				if(SynthButton(event.button.button, false, event, &ev))
					events.push_back(ev);
				break;
			case SDL_ACTIVEEVENT:
				// TODO Shouldn't we pause the game and/or stop rendering on hidden?
				break;
			case SDL_QUIT:
				// TODO Fix this
				// this really isn't what we want to do.
				glopShutDown();
				// Really, we should signal that we have a quit event so the program can deal with it.
				exit(0);
				break;
		}
	}
}

void GlopSetTitle(const string& title) {
	SDL_WM_SetCaption(title.c_str(), NULL);
}

void* GlopCreateWindow(void* title, int x, int y, int width, int height) {
  // this is bad
  if(x == -1) x = 100;
  if(y == -1) y = 100;

  SDL_GL_SetAttribute(SDL_GL_DOUBLEBUFFER, 1);
  SDL_GL_SetAttribute(SDL_GL_RED_SIZE, 1);
  SDL_GL_SetAttribute(SDL_GL_BLUE_SIZE, 1);
  SDL_GL_SetAttribute(SDL_GL_GREEN_SIZE, 1);
  SDL_GL_SetAttribute(SDL_GL_DEPTH_SIZE, 1);
  SDL_GL_SetAttribute(SDL_GL_STENCIL_SIZE, 8);

  SDL_Surface *screen = SDL_SetVideoMode(width, height, 32, SDL_OPENGL);

  GlopSetTitle(string((char*)title));

  return screen;
}

void GlopGetWindowDims(int* x, int* y, int* dx, int* dy) {
	SDL_Surface *screen = SDL_GetVideoSurface();
	*x = 0;
	*y = 0;
	*dx = screen->w;
	*dy = screen->h;
}

// Input functions
// ===============

// See Os.h

static GlopKeyEvent* glop_event_buffer = 0;

void GlopGetInputEvents(void** _events_ret, void* _num_events, void* _horizon) {
  *((long long*)_horizon) = SDL_GetTicks();
  vector<GlopKeyEvent> ret; // weeeeeeeeeeee
  ret.swap(events);

  if (glop_event_buffer != 0) {
    free(glop_event_buffer);
  }

  glop_event_buffer = (GlopKeyEvent*)malloc(sizeof(GlopKeyEvent) * ret.size());
  *((GlopKeyEvent**)_events_ret) = glop_event_buffer;
  *((int*)_num_events) = ret.size();
  for (int i = 0; i < ret.size(); i++) {
    glop_event_buffer[i] = ret[i];
  }
}

void GlopGetMousePosition(int* x, int* y) { // TBI
  *x = 0;
  *y = 0;
}

void GlopSwapBuffers() {
	SDL_GL_SwapBuffers();
}

void GlopEnableVSync(int enable) {
}

} // extern "C"
