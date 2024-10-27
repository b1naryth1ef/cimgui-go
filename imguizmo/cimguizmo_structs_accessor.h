// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

#pragma once

#include "cimguizmo_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void wrap_Style_SetTranslationLineThickness(Style *StylePtr, float v);
extern float wrap_Style_GetTranslationLineThickness(Style *self);
extern void wrap_Style_SetTranslationLineArrowSize(Style *StylePtr, float v);
extern float wrap_Style_GetTranslationLineArrowSize(Style *self);
extern void wrap_Style_SetRotationLineThickness(Style *StylePtr, float v);
extern float wrap_Style_GetRotationLineThickness(Style *self);
extern void wrap_Style_SetRotationOuterLineThickness(Style *StylePtr, float v);
extern float wrap_Style_GetRotationOuterLineThickness(Style *self);
extern void wrap_Style_SetScaleLineThickness(Style *StylePtr, float v);
extern float wrap_Style_GetScaleLineThickness(Style *self);
extern void wrap_Style_SetScaleLineCircleSize(Style *StylePtr, float v);
extern float wrap_Style_GetScaleLineCircleSize(Style *self);
extern void wrap_Style_SetHatchedAxisLineThickness(Style *StylePtr, float v);
extern float wrap_Style_GetHatchedAxisLineThickness(Style *self);
extern void wrap_Style_SetCenterCircleSize(Style *StylePtr, float v);
extern float wrap_Style_GetCenterCircleSize(Style *self);
extern void wrap_Style_SetColors(Style *StylePtr, ImVec4* v);
extern ImVec4* wrap_Style_GetColors(Style *self);
extern ImVec4 cimguizmo_ImVec4_GetAtIdx(ImVec4 *self, int index);

#ifdef __cplusplus
}
#endif
