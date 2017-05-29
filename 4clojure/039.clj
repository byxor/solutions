(fn _interleave [a b]
  (if-not (or (empty? a) (empty? b))
    (concat
      (list (first a) (first b))
      (_interleave (rest a) (rest b)))))
