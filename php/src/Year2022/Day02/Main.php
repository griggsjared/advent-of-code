<?php

declare(strict_types=1);

namespace AOC\Year2022\Day02;

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
        $weaponMap = [
            'X' => 'A',
            'Y' => 'B',
            'Z' => 'C',
        ];

        //foreach round, get the score for the match and then add the weapon score for the self,
        //then sum the scores
        $roundScores = array_map(
            function ($round) use ($weaponMap) {
                $selfWeapon = $weaponMap[$round[1]];
                $opponentWeapon = $round[0];

                $resultScore = $this->getResultScore(
                    $this->getMatchResult($opponentWeapon, $selfWeapon)
                );
                $weaponScore = $this->getWeaponScore($selfWeapon);

                return $resultScore + $weaponScore;
            },
            $this->rounds()
        );

        //sum the scores
        return array_sum($roundScores);
    }

    private function part2(): int
    {
        $possibleWeapons = ['A', 'B', 'C'];

        $matchCache = [];

        //foreach each round, we need to look find the weapon result that === round[1] when we have round[0],
        //then we can get the score for the match and add the weapon score for the self
        $roundScores = array_map(
            function ($round) use ($possibleWeapons, $matchCache) {
                foreach($possibleWeapons as $weapon) {
                    $selfWeapon = $weapon;
                    $opponentWeapon = $round[0];

                    //see if we have already calculated the match result before
                    if(isset($matchCache[$opponentWeapon][$selfWeapon])) {
                        return $matchCache[$opponentWeapon][$selfWeapon];
                    }

                    //calculate the match result
                    $matchResult = $this->getMatchResult($opponentWeapon, $selfWeapon);
                    $matchCache[$opponentWeapon][$selfWeapon] = $matchResult;
                    if($matchResult === $round[1]) {
                        $resultScore = $this->getResultScore($matchResult);
                        $weaponScore = $this->getWeaponScore($selfWeapon);
                        return $resultScore + $weaponScore;
                    }
                }
            },
            $this->rounds()
        );

        return array_sum($roundScores);
    }


    /**
     * @return array<array<string>>
     */
    private function rounds(): array
    {
        //split the input into rounds split on new lines,
        //further split each round into opponent [0] and self [1] split on space
        return array_map(
            fn ($group) => array_map(
                fn ($value) => $value,
                explode(" ", $group)
            ),
            explode("\n", $this->getInput())
        );
    }

    /**
     * X for loss, Y for tie, Z for win
     */
    private function getResultScore(string $result): int
    {
        return match($result) {
            'X' => 0, //loss
            'Y' => 3, //tie
            'Z' => 6, //win
            default => throw new InvalidArgumentException('Invalid result.'),
        };
    }

    /**
     * A = Rock = 1pt
     * B = Paper = 2pt
     * C = Scissors = 3pt
     */
    private function getWeaponScore(string $weapon): int
    {
        return match($weapon) {
            'A' => 1, //rock
            'B' => 2, //paper
            'C' => 3, //scissors
            default => throw new InvalidArgumentException('Invalid weapon.'),
        };
    }

    /**
     * A = Rock
     * B = Paper
     * C = Scissors
     *
     * Z for win, Y for tie, X for loss
     */
    private function getMatchResult(string $opponentWeapon, string $selfWeapon): string
    {
        return match($selfWeapon) {
            'A' => match($opponentWeapon) {
                'C' => 'Z', //rock beats scissors
                'B' => 'X', //rock loses to paper
                'A' => 'Y', //rock ties rock
                default => throw new InvalidArgumentException('Invalid opponent weapon.'),
            },
            'B' => match($opponentWeapon) {
                'A' => 'Z', //paper beats rock
                'C' => 'X', //paper loses to scissors
                'B' => 'Y', //paper ties paper
                default => throw new InvalidArgumentException('Invalid opponent weapon.'),
            },
            'C' => match($opponentWeapon) {
                'B' => 'Z', //scissors beats paper
                'A' => 'X', //scissors loses to rock
                'C' => 'Y', //scissors ties scissors
                default => throw new InvalidArgumentException('Invalid opponent weapon.'),
            },
            default => throw new InvalidArgumentException('Invalid self weapon.'),
        };
    }
}
