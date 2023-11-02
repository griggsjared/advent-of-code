<?php

declare(strict_types=1);

namespace AOC\Year2022\Day04;

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
     * How many of the given pairs contain each other?
     */
    private function part1(): int
    {
        $containsCount = 0;
        foreach($this->pairs() as $pair) {
            if($this->boundsContain($pair[0], $pair[1])) {
                $containsCount++;
            }
        }

        return $containsCount;
    }


    /**
     * How many of the given pairs intersect each other?
     */
    private function part2(): int
    {
        $intersectCount = 0;
        foreach($this->pairs() as $pair) {
            if($this->boundsIntersect($pair[0], $pair[1])) {
                $intersectCount++;
            }
        }

        return $intersectCount;
    }

    private function boundsContain(array $bounds, array $otherBounds): bool
    {
        return ($bounds[0] <= $otherBounds[0] && $bounds[1] >= $otherBounds[1])
            || ($bounds[0] >= $otherBounds[0] && $bounds[1] <= $otherBounds[1]);
    }

    private function boundsIntersect(array $bounds, array $otherBounds): bool
    {
        return ($bounds[0] <= $otherBounds[0] && $bounds[1] >= $otherBounds[0])
            || ($bounds[0] >= $otherBounds[0] && $bounds[0] <= $otherBounds[1]);
    }

    /**
     * @return array<array<array<int>>>
     */
    private function pairs(): array
    {
        //split the input into parts,
        //then split each part split by a comma,
        //then split each part by a dash,
        //then cast each part to an int, the [0] is the lower bound, [1] is the upper bound
        return array_map(
            fn ($group) => array_map(
                fn ($value) => array_map(
                    fn ($part) => (int) $part,
                    explode('-', $value)
                ),
                explode(",", $group)
            ),
            explode("\n", $this->getInput())
        );
    }
}
