<?php

declare(strict_types=1);

namespace AOC\Year2023\Day06;

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

    /**
     * for each item in the array (race), the key is the total allowed time for the race and the value is the min distance need
     * to travel to win, at the start of the race we can wait any number of seconds to increase the travel distance per second,
     * so we can wait 0 seconds and travel 0 distance, or wait 1 second and travel 1 distance, or wait 2 seconds and travel 3 distance. etc.
     * we need to determine the number of combinations for each race to travel to or exceed the distance within the time limit
     * when we wait it counts as a second, so we can't wait more than the time limit
     */
    private function part1(): int
    {
        [$times, $distances] = array_map(
            fn($row) => array_map(
                fn($number) => (int) $number,
                explode(' ' ,preg_replace('/\s+/', ' ', trim($row)))
            ),
            array_map(
                fn($row) => explode(':', $row)[1],
                explode("\n", $this->getInput())
            )
        );

        $races = array_combine($times, $distances); //key = time, value = distance

        $racePossibilities = [];
        foreach($races as $maxTime => $distanceNeeded) {
            $racePossibilities[] = $this->possibleWaitTimesBruteForce($maxTime, $distanceNeeded);
        }

        return array_product($racePossibilities);
    }

    /**
     * Same as above except we remove spaces from the input and use a more efficient algorithm
     */
    private function part2(): int
    {
        [$time, $distance] = array_map(
            fn($row) => (int) preg_replace('/\s+/', '', trim($row)),
            array_map(
                fn($row) => explode(':', $row)[1],
                explode("\n", $this->getInput())
            )
        );

        return $this->possibleWaitTimes($time, $distance);
    }

    /**
     * O(n^2) time complexity and takes forever to run
     */
    private function possibleWaitTimesBruteForce(int $maxTime, int $distanceNeeded): int
    {
        $possibleWaitTimes = 0;

        for($i = 0; $i <= $maxTime; $i++) { //the number of seconds we wait
            $timeLeft = $maxTime - $i;

            //run the race til timer hits the distance
            $timer = 0;
            $distancePerSecond = 1*$i;
            $distanceTraveled = 0;

            while($timer < $timeLeft) {
                $distanceTraveled += $distancePerSecond;
                $timer++;
            }

            if($distanceTraveled > $distanceNeeded) {
                $possibleWaitTimes++;
            }
        }

        return $possibleWaitTimes;
    }

    /**
     * O(n) time complexity
     */
    private function possibleWaitTimes(int $maxTime, int $distanceNeeded): int
    {
        $possibleWaitTimes = 0;

        for($i = 1; $i <= $maxTime; $i++) { //the number of seconds we wait
            $timeLeft = $maxTime - $i;

            //calculate the distance directly
            $distancePerSecond = $i;
            $distanceTraveled = $distancePerSecond * $timeLeft;

            if($distanceTraveled >= $distanceNeeded) {
                $possibleWaitTimes++;
            }
        }

        return $possibleWaitTimes;
    }
}
