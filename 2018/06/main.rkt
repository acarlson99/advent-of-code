#lang racket

(define inp
  (map (compose
        (λ (l) (cons (car l) (cadr l)))            ; make pairs
        (λ (l) (map (compose car string->list) l)) ; strings to chars
        cdr
        (λ (s) (regexp-match #px"Step (.+) must be finished before step (.+) can begin\\." s)))
       (file->lines "input.txt")))
;; POCUEFIXHRGWDZABTQJYMNKVSL

(define parents (make-hash
                 (map
                  (λ (l) (cons l null))
                  (map car inp))))

(define (add-to-map m k v)
  (hash-set! m k
             (cons v (hash-ref m k null))))

(for ([p (in-list inp)])
  (add-to-map parents (cdr p) (car p)))

(define children (make-hash
                  (map
                   (λ (l) (cons l null))
                   (map car inp))))

(for ([p (in-list inp)])
  (add-to-map children (car p) (cdr p)))

(define (solve-1)
  (define visited-hash (make-hash (map (λ (c) (cons c #f)) (string->list "ABCDEFGHIJKLMNOPQRSTUVWXYZ"))))

  (define (visited? v)
    (hash-ref visited-hash v #f))

  (define (visit v)
    (hash-set! visited-hash v #t))

  (define (visitable v)
    (foldr (λ (a b) (and a b)) #t (map (λ (k) (hash-ref visited-hash k)) (hash-ref parents v '()))))

  (define steps
    (sort
     (foldr (λ (x acc) (or (and x (cons x acc)) acc))
            '()
            (hash-map parents (λ (k v) (and (null? v) k))))
     char<?))

  (define visit-order '())

  (for ([n (in-naturals)])
    #:break (empty? steps)
    (let ([v (car steps)])
      (set! steps (cdr steps))
      (if (and (not (visited? v)) (visitable v))
          (begin
            (visit v)
            (set! visit-order (cons v visit-order))
            ;; ref `parents`, add to list of steps
            (set! steps
                  (sort
                   (append steps (hash-ref children v null))
                   char<?)))
          null)))

  (displayln (list->string (reverse visit-order))))

(solve-1)
