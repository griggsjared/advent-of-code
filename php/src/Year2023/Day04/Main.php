<?php

declare(strict_types=1);

namespace AOC\Year2023\Day04;

use AOC\MainInterface;
use AOC\MainTrait;
use AOC\PartEnum;
use InvalidArgumentException;

class Main implements MainInterface
{
    use MainTrait;

    public function run(PartEnum $part): int
    {
        return match($part) {
            PartEnum::P1 => $this->part1(),
            PartEnum::P2 => $this->part2(),
            default => throw new InvalidArgumentException('Invalid part.'),
        };
    }

    private function part1(): int
    {
        //for every row, parse into winners and picks,
        //the first pick in the winners counts as 1 point, each additional pick doubles the points,
        $cardPoints = array_map(
            function($count) {
                // 1 = 1, 2 = 2, 3 = 4, 4 = 8, 5 = 16
                return $count > 0 ? 2 ** ($count - 1) : 0;
            },
            $this->getCardWinnersCounts()
        );

        return array_sum($cardPoints);
    }

    private function part2(): int
    {
        $winnerCounts = $this->getCardWinnersCounts();

        //the number of each card we get, starts as 0,
        $cardCounts = array_map(
            fn() => 1,
            $winnerCounts
        );

        //for each card, depending on the number of winners, we get a number of card from the subsequent cards
        foreach($winnerCounts as $index => $winnerCount) {
            $currentCardCount = $cardCounts[$index];
            for($i = 1; $i <= $winnerCount; $i++) {
                if(!isset($cardCounts[$index + $i])) {
                    break;
                }
                //add the current card count to the subsequent cards
                $cardCounts[$index + $i] += $currentCardCount;
            }
        }

        return array_sum($cardCounts);
    }

    private function getCardWinnersCounts(): array
    {
        return array_map(
            function($row) {

                //parsing each row into winners and picks
                [$winners, $picks] = array_map(
                    fn($part) => array_map(
                        fn($number) => (int) $number,
                        explode(' ' ,preg_replace('/\s+/', ' ', trim($part)))
                    ),
                    //discarding the cart number and splitting the winners and picks
                    explode('|', explode(':',$row)[1])
                );

                $winnersCount = 0;
                foreach($picks as $pick) {
                    if(in_array($pick, $winners)) {
                        $winnersCount++;
                    }
                }

                return $winnersCount;
            },
            explode("\n", $this->getInput())
        );
    }
}
