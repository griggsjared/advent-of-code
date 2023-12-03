<?php

declare(strict_types=1);

namespace AOC\Year2023\Day02;

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
        $maxCubes = [
            'red' => 12,
            'green' => 13,
            'blue' => 14
        ];

        $validGameIds = array_map(
            function($game) use ($maxCubes) {
                $return = (int) $game['id']; //score for game is the id, but if invalid, score is 0
                foreach($game['rounds'] as $round) {
                    foreach($round as $cube) {
                        if($cube['amount'] > $maxCubes[$cube['color']]) {
                            $return = 0; //invalid game
                            break;
                        }
                    }
                }
                return $return;
            },
            $this->getParsedGames()
        );

        return array_sum($validGameIds);
    }

    private function part2(): int
    {

        $roundPowers = array_map(
            function($game) {
                //for each game what is the min of each color needed to make the game (all rounds) valid
                $maxCubes = [
                    'red' => 0,
                    'green' => 0,
                    'blue' => 0
                ];

                foreach($game['rounds'] as $round) {
                    foreach($round as $cube) {
                        if($cube['amount'] > $maxCubes[$cube['color']]) {
                            $maxCubes[$cube['color']] = $cube['amount'];
                        }
                    }
                }

                //return the power, which is the min number of each multiplied together
                return array_product($maxCubes);
            },
            $this->getParsedGames()
        );

        return array_sum($roundPowers);
    }

    private function getParsedGames(): array
    {
        return array_map(
            function($game) {
                $idToRounds = explode(':', str_replace('Game ', '', $game));
                $id = $idToRounds[0];

                $rounds = array_map(
                    fn($round) => array_map(
                        fn($cube) => explode(' ', trim($cube)),
                        $round
                    ),
                    array_map(
                        fn($cube) => explode(',', $cube),
                        explode(';', $idToRounds[1])
                    )
                );

                $rounds = array_map(
                    fn($round) => array_map(
                        fn($cube) => [
                            'color' => $cube[1],
                            'amount' => (int) $cube[0]
                        ],
                        $round
                    ),
                    $rounds
                );

                return [
                    'id' => $id,
                    'rounds' => $rounds
                ];
            },
            explode("\n", $this->getInput())
        );
    }

}
