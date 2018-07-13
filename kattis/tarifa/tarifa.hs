import Control.Monad

readAnInt :: IO Int
readAnInt = readLn

moneyNextMonth limit months spendings =
    (limit * (months + 1)) - (sum spendings)

main = do
    limit <- readAnInt
    months <- readAnInt
    spendings <- replicateM months readAnInt

    let answer = moneyNextMonth limit months spendings
    putStrLn (show answer)