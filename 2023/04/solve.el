;;;                                                  -*- lexical-binding: t; -*-

(defun f (ls)
  (let ((id (car ls))
		(winners (cadr ls))
		(gottem (caddr ls)))
	(apply #'append
		   (mapcar
			(lambda (x) (and (memq x gottem) (list x)))
			winners))))

(defun aggregate (ls)
  (let
	  ((len (length (f ls))))
	len))

;; prolly wanna run `sh make-lisp.sh < input.txt > input.el' there chief
(let* ((input__ '(
				  (1 (41 48 83 86 17) (83 86  6 31 17  9 48 53))
				  (2 (13 32 20 16 61) (61 30 68 82 17 32 24 19))
				  (3 (1 21 53 59 44) (69 82 63 72 16 21 14  1))
				  (4 (41 92 73 84 69) (59 84 76 51 58  5 54 83))
				  (5 (87 83 26 28 32) (88 30 70 12 93 22 82 36))
				  (6 (31 18 13 56 72) (74 77 10 23 35 67 36 11))
				  ))
	   (lens (mapcar #'aggregate input)))
										; P1
  (message "P1-- %d" (apply #'+ (mapcar
								 (lambda (n)
								   (if (> n 0)
									   (expt 2 (- n 1))
									 0))
								 lens)))
										; P2
  (message "P2-- %d" (let ((cards '()))
					   (mapcar
						(lambda (n)
						  (let* ((new-cards (cons 1 (take n cards)))
								 (card-count (apply #'+ new-cards)))
							(setq cards (cons card-count cards))))
						(reverse lens))
					   (apply #'+ cards))))
