#!emacs --script
;; -*- lexical-binding: t; -*-

;; AOC 2023 day 06

(require 'cl-lib)

(defun f (n) `(lambda (bt)  (* (- ,n bt) bt)))

;; old solution; not good
;; (apply #'*
;; 	   (let ((duration-dists
;; 			  (mapcar* #'cons
;; 					   '(    34     90     89     86) ; duration
;; 					   '(   204   1713   1210   1780) ; best distance
;; 					   )))
;; 		 (mapcar
;; 		  (lambda (pair)
;; 			(let* ((duration (car pair))
;; 				   (distance (cdr pair))
;; 				   (fn (f duration)))
;; 			  (length
;; 			   (seq-filter (lambda (n) (>= n distance))
;; 						   (mapcar fn (number-sequence 0 duration))))))
;; 		  duration-dists)))

;; new solution; much better
(defun quadratic (a b c)
  (list (/
		 (+ (- b) (sqrt (- (expt b 2.0) (* 4 a c))))
		 (* 2 a))
		(/
		 (- (- b) (sqrt (- (expt b 2.0) (* 4 a c))))
		 (* 2 a))))

(defun min-max-quadratic-difference (n best)
  (let* ((qs (quadratic -1.0 n (- best)))
		(bt-max (floor (apply #'max qs)))
		(bt-min (ceiling (apply #'min qs))))
	(+ 1 (- bt-max bt-min))))

;; p1

(defvar durations '(    34     90     89     86)) ; duration
(defvar distances '(   204   1713   1210   1780)) ; best distance

(defun p1 ()
  (let* ((duration-dists
		  (cl-mapcar #'cons durations distances))
		 (button-ranges
		  (mapcar
		   (lambda (pair)
			 (min-max-quadratic-difference (car pair) (cdr pair)))
		   duration-dists)))
	(apply #'* button-ranges)))

(message "P1-- %d" (p1))

;; p2

;;         duration     best distance
;; -1t^2 + 34908986t - 204171312101780 = 0

(defvar p2-duration
  (+ 0.0 (string-to-number (mapconcat #'int-to-string durations))))
(defvar p2-distance
  (+ 0.0 (string-to-number (mapconcat #'int-to-string distances))))

(defun aoc-p2-test ()
  (let* ((n p2-duration)
		 (best p2-distance)
		 (qs (quadratic -1.0 n (- best)))
		 (bt-max (floor (apply #'max qs)))
		 (bt-min (ceiling (apply #'min qs))))
	(and
	 (<=
	  best
	  (funcall (f n) bt-max))
	 (not (<=
		   best
		   (funcall (f n) (+ bt-max 1))))
	 (<=
	  best
	  (funcall (f n) bt-min))
	 (not (<=
		   best
		   (funcall (f n) (- bt-min 1))))
	 (+ 1 (- bt-max bt-min)))))

(defun p2 ()
  (min-max-quadratic-difference p2-duration p2-distance))

(message "P2-- %d" (p2))
