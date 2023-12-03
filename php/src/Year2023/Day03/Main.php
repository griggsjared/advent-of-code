<?php

declare(strict_types=1);

namespace AOC\Year2023\Day03;

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
     * Find all symbols that are not numbers and are not dots.
     * For each symbol, check all neighbors for numbers.
     * If a number is found, add it to the valid numbers map.
     * Sum all the valid numbers and return the sum.
     * Numbers cannot be used twice to storing the address of the number in the map.
     */
    private function part1(): int
    {
        $rows = explode("\n", $this->getInput());

        $validNumbersMap = [];

        $columnsCount = strlen($rows[0]);
        $rowCount = count($rows);

        foreach($rows as $rowNumber => $row) {

            for($columnNumber = 0; $columnNumber < $columnsCount; $columnNumber++) {

                $character = $row[$columnNumber]; // the character at the current position

                if($character === '.' || is_numeric($character)) {
                    continue;
                }

                foreach($this->getNeighborsMap() as $neighbor) {
                    $neighborRow = $rowNumber + $neighbor[0];
                    $neighborColumn = $columnNumber + $neighbor[1];

                    $numberData = $this->checkPositionForNumber($rows, $neighborRow, $neighborColumn);
                    if($numberData) {
                        $validNumbersMap[$numberData['address']] =  $numberData['number'];
                    }
                }
            }
        }

        return array_sum($validNumbersMap);
    }

    /**
     * Find the gear ratios of all the gears in the grid.
     * If the gear has exactly two neighbors, then the gear ratio is the product of the two neighbors.
     * We sum all the gear ratios and return the sum.
     */
    public function part2(): int
    {
        $rows = explode("\n", $this->getInput());

        $gearRatios = [];

        foreach($rows as $rowNumber => $row) {

            for($columnNumber = 0; $columnNumber < strlen($rows[0]); $columnNumber++) {

                $character = $row[$columnNumber];

                if($character !== '*') {
                    continue;
                }

                $validNumbersForCharacter = [];

                foreach ($this->getNeighborsMap() as $neighbor) {
                    $neighborRow = $rowNumber + $neighbor[0];
                    $neighborColumn = $columnNumber + $neighbor[1];

                    $numberData = $this->checkPositionForNumber($rows, $neighborRow, $neighborColumn);
                    if ($numberData) {
                        $validNumbersForCharacter[$numberData['address']] = $numberData['number'];
                    }
                }

                if(count($validNumbersForCharacter) === 2) {
                    $gearRatios[] = array_product($validNumbersForCharacter);
                }
            }
        }

        return array_sum($gearRatios);
    }

    private function checkPositionForNumber(array $rows, int $rowPosition, int $columnPosition): ?array
    {
        //make sure we are not out of bounds
        if($rowPosition < 0 || $rowPosition >= count($rows)) {
            return null;
        }

        $row = $rows[$rowPosition];
        $character = $row[$columnPosition];

        if(is_numeric($character)) {

            $fullNumber = $character;

            $columnsCount = strlen($row);

            //check to the left until we find a non number or the start of the row
            $leftColumnPosition = $columnPosition;
            while($leftColumnPosition > 0) {
                $leftColumnPosition--;
                $leftCharacter = $row[$leftColumnPosition];
                if(is_numeric($leftCharacter)) {
                    $fullNumber = $leftCharacter . $fullNumber;
                } else {
                    break;
                }
            }

            //check to the right until we find a non number or the end of the row
            $rightColumnPosition = $columnPosition;
            while($rightColumnPosition < $columnsCount - 1) {
                $rightColumnPosition++;
                $rightCharacter = $row[$rightColumnPosition];
                if(is_numeric($rightCharacter)) {
                    $fullNumber = $fullNumber . $rightCharacter;
                } else {
                    break;
                }
            }

            return [
                'number' => (int) $fullNumber,
                'address' => $rowPosition . $leftColumnPosition . $rightColumnPosition
            ];
        }

        return null;
    }

    /**
     * @return array<int, array<int, int>> 0 = row, 1 = column
     */
    private function getNeighborsMap(): array
    {
        return [
            [-1, -1], // top left
            [-1, 0], // top
            [-1, 1], // top right
            [0, -1], // left
            [0, 1], // right
            [1, -1], // bottom left
            [1, 0], // bottom
            [1, 1] // bottom right
        ];
    }
}
