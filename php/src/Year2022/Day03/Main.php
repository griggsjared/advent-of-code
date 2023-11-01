<?php

declare(strict_types=1);

namespace AOC\Year2022\Day03;

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
        /**
         * @var array<array<string>>
         */
        $splitBags = array_map(
            function ($bag) {
                $length = strlen($bag);
                $half = $length / 2;
                return [
                    substr($bag, 0, $half),
                    substr($bag, $half, $length)
                ];
            },
            explode("\n", $this->getInput())
        );

        //a single character will appear in each compartment of each bag, find that character in each bag and get the priority value
        $bagsValues = array_map(
            function ($bag) {
                $compartment1 = array_flip(str_split($bag[0]));
                $compartment2 = str_split($bag[1]);
                $foundChar = null;
                foreach($compartment2 as $char) {
                    if (isset($compartment1[$char])) {
                        $foundChar = $char;
                        break;
                    }
                }
                return $this->priorityMap()[$foundChar];
            },
            $splitBags
        );

        return array_sum($bagsValues);
    }

    private function part2(): int
    {
        /**
         * @var array<array<string>>
         */
        $groupBags = array_chunk(
            explode("\n", $this->getInput()),
            3
        );

        //a single character will appear in all 3 bags, find the character in each bag and get the priority value
        $groupValues = array_map(
            function ($group) {
                $bag1 = array_flip(str_split($group[0]));
                $bag2 = array_flip(str_split($group[1]));
                $bag3 = str_split($group[2]);
                $foundChar = null;
                foreach($bag3 as $char) {
                    if (isset($bag1[$char]) && isset($bag2[$char])) {
                        $foundChar = $char;
                        break;
                    }
                }
                return $this->priorityMap()[$foundChar];
            },
            $groupBags
        );

        return array_sum($groupValues);
    }

    /**
     * @return array<string, int>
     */
    private function priorityMap(): array
    {
        $map = array_merge(range('a', 'z'), range('A', 'Z'));

        return array_combine(
            $map,
            array_map(
                fn ($value) => $value + 1,
                array_keys($map)
            )
        );
    }
}
