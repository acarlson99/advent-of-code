#lang racket

(define inp (file->lines "input.txt"))

(define (buccet l)
  (cond
    [(empty? l) '()]
    [(empty? (cdr l)) (list (list (car l) 1))]
    [else (let ([c (car l)]
                [nc (cadr l)]
                [rst (buccet (cdr l))])
            (if (eqv? c nc)
                (cons
                 (list c (+ 1 (cadar rst)))
                 (cdr rst))
                (cons (list c 1) rst)))]))

;; count frequency of chars in string
(define (count-chars s)
  (buccet (sort (string->list s) char<?)))

(define (list->freq-map lst [mp (make-hash)])
  (for ([l (in-list lst)])
    (hash-update! mp l (λ (n) (+ n 1)) 0))
  mp)

;; find number of 2,3 instances
(define (map-chk mp)
  (define (to-num n)
    (if n
        1
        0))
  (list (to-num (hash-ref! mp 2 #f)) (to-num (hash-ref! mp 3 #f))))

(define (condense a b)
  (for/list ([na (in-list a)]
             [nb (in-list b)])
    (+ na nb)))

(displayln (apply * (foldr condense '(0 0)
                        (map
                         (compose map-chk list->freq-map (λ (a) (map cadr a)) count-chars)
                         inp))))

(define (count-list-diff s1 s2)
  (cond
    [(empty? s1) (length s2)]
    [(empty? s2) (length s1)]
    [else (+
           (count-list-diff (cdr s1) (cdr s2)) (if (eqv? (car s1) (car s2))
                                                   0
                                                   1))]))

(define (trim-list-diff a b)
  (cond
    [(empty? a) b]
    [(empty? b) a]
    [else
     (append
      (if (eqv? (car a) (car b))
          (list (car a))
          '())
      (trim-list-diff (cdr a) (cdr b)))]))

(define done #f)

(for ([v1 (in-list (map string->list inp))]
      #:break done)
  (for ([v2 (in-list (map string->list inp))]
        #:break done)
    (and (eqv? (count-list-diff v1 v2) 1)
         (and (set! done #t) (displayln (list->string (trim-list-diff v1 v2)))))))
