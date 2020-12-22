#lang racket

(require math)

;; #7 @ 129,812: 18x17
(define inp
  (map
   (λ (s) (cdr (regexp-match #px"(.*) @ (\\d+),(\\d+): (\\d+)x(\\d+)" s)))
   (file->lines "input.txt")))

;; part one
(define arr (array->mutable-array (make-array #[1000 1000] 0)))

(define (fill-arr x y xl yl)
  (for ([xr (in-range xl)])
    (for ([yr (in-range yl)])
      (let ([xp (+ x xr)]
            [yp (+ y yr)])
        (array-set! arr `#(,xp ,yp) (+ 1 (array-ref arr `#(,xp ,yp))))))))

(for ([n (in-list inp)])
  (match n
    [(list name x y xl yl)
     (fill-arr (string->number x) (string->number y) (string->number xl) (string->number yl))]))

(displayln (foldr (lambda (a b) (if (>= a 2) (+ 1 b) b)) 0 (array->list arr)))

;; part two
(define arr-2 (array->mutable-array (make-array #[1000 1000] 0)))

(define hs (make-hash (map (lambda (n) (cons (car n) #t)) inp)))

(define (remove-from-map name)
  (hash-set! hs name #f))

(define (fill-arr-2 name x y xl yl)
  (for ([xr (in-range xl)])
    (for ([yr (in-range yl)])
      (let ([xp (+ x xr)]
            [yp (+ y yr)])
        (define point `#(,xp ,yp))
        (define v (array-ref arr-2 point))
        (cond
          [(eqv? v 0) (array-set! arr-2 point name)] ; no overlap (yet)
          [(eqv? v -1) (remove-from-map name)] ; old overlap
          [else
           (begin
             (remove-from-map name)
             (remove-from-map v)
             (array-set! arr-2 point -1))]))))) ; new overlap

(for ([n (in-list inp)])
  (match n
    [(list name x y xl yl)
     (fill-arr-2 name
                 (string->number x) (string->number y) (string->number xl) (string->number yl))]))

(displayln (substring (foldr (lambda (a b) (if (cdr a) (car a) b)) 0 (hash->list hs)) 1))
