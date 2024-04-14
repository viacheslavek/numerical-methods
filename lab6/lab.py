from sympy import symbols, diff, atan
import math

eps = 0.001


def f(x, y):
    return (x ** 2) + (2 * y ** 2) + (2 * x) + 0.3 * atan(x * y)


def analytic_min():
    return -1.01124535766275, 0.0754049591595788


def find_min(x, y):
    fx = diff(f(x, y), x)
    fy = diff(f(x, y), y)

    print('df/dx: ', fx)
    print('df/dy: ', fy)

    fxx = diff(fx, x)
    fxy = diff(fx, y)
    fyy = diff(fy, y)

    print('df2/dx2: ', fxx)
    print('df2/dxdy: ', fxy)
    print('df2/dy2: ', fyy)

    print()

    xk, yk = 0.0, 1.0

    counter = 0

    while max(fx.subs({x: xk, y: yk}), fy.subs({x: xk, y: yk})) >= eps:

        fx_value = fx.subs({x: xk, y: yk})
        fy_value = fy.subs({x: xk, y: yk})

        fxx_value = fxx.subs({x: xk, y: yk})
        fxy_value = fxy.subs({x: xk, y: yk})
        fyy_value = fyy.subs({x: xk, y: yk})

        phi = -fx_value**2 - fy_value**2
        eta = fxx_value * fx_value**2 + 2 * fxy_value * fx_value * fy_value + fyy_value * fy_value**2

        t = - phi / eta

        xk = xk - t * fx_value
        yk = yk - t * fy_value

        counter += 1

    print(f'find for {counter} iterate')
    print()

    return xk, yk


def main():
    print('f: ', f(symbols('x'), symbols('y')))

    x, y = find_min(symbols('x'), symbols('y'))

    print(f'methods min: {x, y}')
    print(f'original min: {analytic_min()}')
    print(f'difference: {math.fabs(x - analytic_min()[0]), math.fabs(y - analytic_min()[1])}')


main()
