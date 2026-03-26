#include "calc.h"

/* Таблиця цін за 1 см²: [тип скла 0/1][матеріал рами 0..2] */
static const double prices[2][3] = {
    {2.5, 0.5, 1.5}, /* однокамерний */
    {3.0, 1.0, 2.0}  /* двокамерний */
};

static const double sill_price = 350.0;

/*
 * calc_glass_price — вартість склопакета (грн).
 * Площа = width * height (см²), множиться на ціну за см².
 * hasSill != 0 — додається вартість підвіконня.
 */
double calc_glass_price(double width, double height, int glassType, int frameType, int hasSill) {
    if (glassType < 0 || glassType > 1 || frameType < 0 || frameType > 2) {
        return 0.0;
    }
    double total = width * height * prices[glassType][frameType];
    if (hasSill) {
        total += sill_price;
    }
    return total;
}
