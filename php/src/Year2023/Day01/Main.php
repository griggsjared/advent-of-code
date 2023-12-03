<?php

declare(strict_types=1);

namespace AOC\Year2023\Day01;

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
        return $this->sumForRows(explode("\n", $this->getInput()));
    }

    private function part2(): int
    {
        $wordNumberMap = [
            'one' => '1',
            'two' => '2',
            'three' => '3',
            'four' => '4',
            'five' => '5',
            'six' => '6',
            'seven' => '7',
            'eight' => '8',
            'nine' => '9',
            '1' => '1',
            '2' => '2',
            '3' => '3',
            '4' => '4',
            '5' => '5',
            '6' => '6',
            '7' => '7',
            '8' => '8',
            '9' => '9'
        ];

        //manual search
        $rows = array_map(
            function($row) use ($wordNumberMap) {
                //searching for substrings..
                //starting at the start of the row,
                //search for a word matching any of the keys in the map, if found, add the mapped value to a new row, then move to the next character and repeat
                $newRow = '';
                for($i = 0; $i < strlen($row); $i++) {
                    $substring = substr($row, $i);
                    foreach($wordNumberMap as $word => $number) {
                        if(str_starts_with($substring, (string) $word)) {
                            $newRow .= $number;
                            break;
                        }
                    }
                }

                return $newRow;
            },
            explode("\n", $this->getInput())
        );

        return $this->sumForRows($rows);
    }

    private function sumForRows(array $rows): int
    {
        //split each row into an array of characters, then filter out any non-numeric characters
        $numbers = array_map(
            fn($row) => array_values(
                array_filter(
                    str_split($row),
                    fn($character) => is_numeric($character)
                )
            ),
            $rows
        );

        //map each row to an array with the first and last number
        $firstAndLast = array_map(
            fn($row) => [
                'first' => $row[0],
                'last' => $row[count($row) - 1]
            ],
            $numbers
        );

        $combined = array_map(
            function($row) {
                return (int) ($row['first'] . $row['last']);
            },
            $firstAndLast
        );

        return array_sum($combined);
    }
}
