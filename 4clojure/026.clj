(fn fibonacci_sequence [n]
  (if (= n 2)
    [1 1]
    (let [sequence (fibonacci_sequence (- n 1))]
      (conj
        sequence
        (+ (last sequence) (nth sequence (- (count sequence) 2)))))))
