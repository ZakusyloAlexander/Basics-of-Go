#ifndef CALC_H
#define CALC_H

/*
 * Оголошення функції, яку викликає Go через cgo.
 * glassType: 0 — однокамерний, 1 — двокамерний
 * frameType: 0 — дерев'яний, 1 — металевий, 2 — металопластиковий
 * hasSill: 1 — з підвіконням, 0 — без
 */
double calc_glass_price(double width, double height, int glassType, int frameType, int hasSill);

#endif
