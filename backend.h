#pragma once

#include "cimgui_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef int GLFWWindowFlags;
enum GLFWWindowFlags_ {
  GLFWWindowFlagsNone = 0,
  GLFWWindowFlagsNotResizable = 1 << 0,
  GLFWWindowFlagsMaximized = 1 << 1,
  GLFWWindowFlagsFloating = 1 << 2,
  GLFWWindowFlagsFrameless = 1 << 3,
  GLFWWindowFlagsTransparent = 1 << 4,
};

typedef struct GLFWwindow GLFWwindow;
typedef struct GLFWmonitor GLFWmonitor;
struct GLFWwindow;
struct GLFWmonitor;

typedef void (*CloseCallback)();
typedef void (*DropCallback)(int path_count, const char* paths[]);
typedef void (*VoidCallback)();

extern void igSetBgColor(ImVec4 color);
extern void igSetTargetFPS(unsigned int fps);
extern GLFWwindow *igCreateGLFWWindow(const char *title, int width, int height, GLFWWindowFlags flags,
                                      VoidCallback afterCreateContext);
extern void igRunLoop(GLFWwindow *window, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
                      VoidCallback beforeDestroyContext);
extern void igGLFWWindow_GetDisplaySize(GLFWwindow *window, int *width, int *height);
extern void igGLFWWindow_SetShouldClose(GLFWwindow *window, int value);
extern void igRefresh();
extern ImTextureID igCreateTexture(unsigned char *pixels, int width, int height);
extern void igDeleteTexture(ImTextureID id);
extern void glfw_set_drop_callback(GLFWwindow *window, DropCallback callback);
extern void glfw_hide_window(GLFWwindow *window);
extern void glfw_show_window(GLFWwindow *window);
extern void glfw_set_closable(bool value);
extern void glfw_set_close_callback(GLFWwindow* window, CloseCallback callback);

#ifdef __cplusplus
}
#endif
