import Data.List (delete)

findCharactersToRemove :: String -> String -> [Int]
findCharactersToRemove first second =
  [i + 1 | i <- [0..length first - 1], let removedString = take i first ++ drop (i + 1) first, removedString == second]

main :: IO ()
main = do
  first <- getLine
  second <- getLine
  if length first > length second + 1
    then putStrLn "0"
    else do
      let indexesToRemove = findCharactersToRemove first second
      putStrLn $ show (length indexesToRemove)
      if not (null indexesToRemove)
        then putStrLn $ unwords (map show indexesToRemove)
        else return ()
