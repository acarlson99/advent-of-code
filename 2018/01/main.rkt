#lang racket

(define inp (map string->number (file->lines "input.txt")))

(displayln (apply + (map string->number (file->lines "input.txt"))))

(define total 0)
(define m (make-hash))
(define done #f)

(for ([i (in-naturals)]
      #:break done)
  (for ([n (in-list inp)]
        #:break done)
    (if (hash-ref m total #f)
        (begin
          (displayln total)
          (set! done #t))
        (begin
          (hash-set! m total #t)
          (set! total (+ total n))))))
