<?php

declare(strict_types=1);

namespace AOC\Year2022\Day01;

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
        return max($this->groupsSums());
    }

    private function part2(): int
    {
        $groupsSumsRanked = $this->groupsSums();

        //sort the sums in descending order
        rsort($groupsSumsRanked);

        //get the top 3
        $top3 = array_slice($groupsSumsRanked, 0, 3);

        return array_sum($top3);
    }

    /**
     * @return array<array<int>>
     */
    private function groups(): array
    {
        //split into groups and values for each row in group, also cast to int
        return array_map(
            fn($group) => array_map(
                fn($value) => (int) $value,
                explode("\n", $group)
            ),
            explode("\n\n", $this->getInput())
        );
    }

    /**
     * @return array<int>
     */
    private function groupsSums(): array
    {
        //sum for each group
        return array_map(
            fn($group) => array_sum($group),
            $this->groups()
        );
    }
}
