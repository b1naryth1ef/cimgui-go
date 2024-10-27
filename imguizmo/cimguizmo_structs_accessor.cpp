// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.


#include <string.h>
#include "cimguizmo_wrapper.h"
#include "cimguizmo_structs_accessor.h"

void wrap_Style_SetTranslationLineThickness(Style *StylePtr, float v) { StylePtr->TranslationLineThickness = v; }
float wrap_Style_GetTranslationLineThickness(Style *self) { return self->TranslationLineThickness; }
void wrap_Style_SetTranslationLineArrowSize(Style *StylePtr, float v) { StylePtr->TranslationLineArrowSize = v; }
float wrap_Style_GetTranslationLineArrowSize(Style *self) { return self->TranslationLineArrowSize; }
void wrap_Style_SetRotationLineThickness(Style *StylePtr, float v) { StylePtr->RotationLineThickness = v; }
float wrap_Style_GetRotationLineThickness(Style *self) { return self->RotationLineThickness; }
void wrap_Style_SetRotationOuterLineThickness(Style *StylePtr, float v) { StylePtr->RotationOuterLineThickness = v; }
float wrap_Style_GetRotationOuterLineThickness(Style *self) { return self->RotationOuterLineThickness; }
void wrap_Style_SetScaleLineThickness(Style *StylePtr, float v) { StylePtr->ScaleLineThickness = v; }
float wrap_Style_GetScaleLineThickness(Style *self) { return self->ScaleLineThickness; }
void wrap_Style_SetScaleLineCircleSize(Style *StylePtr, float v) { StylePtr->ScaleLineCircleSize = v; }
float wrap_Style_GetScaleLineCircleSize(Style *self) { return self->ScaleLineCircleSize; }
void wrap_Style_SetHatchedAxisLineThickness(Style *StylePtr, float v) { StylePtr->HatchedAxisLineThickness = v; }
float wrap_Style_GetHatchedAxisLineThickness(Style *self) { return self->HatchedAxisLineThickness; }
void wrap_Style_SetCenterCircleSize(Style *StylePtr, float v) { StylePtr->CenterCircleSize = v; }
float wrap_Style_GetCenterCircleSize(Style *self) { return self->CenterCircleSize; }
void wrap_Style_SetColors(Style *StylePtr, ImVec4* v) { memcpy(StylePtr->Colors, v, sizeof(ImVec4)*15); }
ImVec4* wrap_Style_GetColors(Style *self) { return self->Colors; }
ImVec4 cimguizmo_ImVec4_GetAtIdx(ImVec4 *self, int index) { return self[index]; }
