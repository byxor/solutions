(fn _interpose [value _sequence]
  (drop-last
    ((fn __interpose [value _sequence]
      (if-not (empty? _sequence)
        (concat
          (list (first _sequence))
          (list value)
          (__interpose value (rest _sequence)))))
      value
      _sequence)))
