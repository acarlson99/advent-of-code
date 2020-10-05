#lang racket

(define inp (string->list (car (file->lines "input.txt"))))

(define (reduce lst)
  (match lst
    [(list a b c ...)
     #:when (and
             (eqv? (char-downcase a) (char-downcase b))
             (not (eqv? a b)))
     (reduce c)]
    [(list a b c ...) (cons a (reduce (cons b  c)))]
    [else lst]))

(define (reduce-full lst)
  (let ([old lst]
        [new (reduce lst)])
    (for ([n (in-naturals)])
      #:break (equal? old new)
      (begin
        (set! old new)
        (set! new (reduce new))
        ))
    old))

(define reduced-inp (reduce-full inp))

(displayln (length reduced-inp))

(define (remove-all c lst)
  (match lst
    [(list a b ...)
     #:when (equal? (char-downcase a) (char-downcase c))
     (remove-all c b)]
    [(list a b ...)
     (cons a (remove-all c b))]
    [else lst]))

(define (minim lst)
  (foldl 
   (lambda (x y) (if (< x y) x y))
   (first lst)
   (rest lst)))

(displayln
 (apply min
        (map
         (λ (x) (length (reduce-full (remove-all x reduced-inp))))
         (string->list "abcdefghijklmnopqrstuvwxyz"))))
