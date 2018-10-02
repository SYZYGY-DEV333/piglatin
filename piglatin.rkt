(define lst '("sh" "gl" "ch" "ph" "tr" "br" "fr" "bl" "gr" "st" "sl" "sl" "pl" "fl" "th"))

(define (flatten list)
    (cond ((null? list) '())
        [(list? (car list)) (append (flatten (car list)) (flatten (cdr list)))]
        [else (cons (car list) (flatten (cdr list)))]))

(define (member? str list) (if [member str list] #t #f))

(define (piglatinize word)
    (let [(i (string->list word))]
        (cond 
            [(member? (car i) '(#\a #\e #\i #\o #\u))
                (string-append word "ay")]
            [(member? (string-append (list->string (flatten (take i 2)))) lst)
                (string-append (list->string (flatten (drop i 2))) (list->string (flatten (take i 2))) "ay")]
            [(not (regexp-match #rx"^[A-Za-z]+$" word))
                word]
            [else (string-append (list->string (flatten (list (drop i 1) (car i)))) "ay")])))

(define (piglatin phrase)
    (string-join (map piglatinize (string-split phrase))))

(define (main)
    (display "Type what you would like translated into pig-latin and press ENTER: ")
    (displayln (piglatin (read-line))))

(main)
