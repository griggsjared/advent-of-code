<?php

declare(strict_types=1);

namespace AOC;

use InvalidArgumentException;
use ReflectionClass;

trait MainTrait
{
    private string $input;

    public function __construct()
    {
        $dir = dirname((new ReflectionClass($this))->getFilename());

        if (file_exists($dir  . '/input.txt')) {
            $this->setInput(file_get_contents($dir . '/input.txt'));
        }
    }

    public function setInput(string $string): void
    {
        $this->input = trim($string);
    }

    public function getInput(): string
    {
        if(!isset($this->input)) {
            throw new InvalidArgumentException(
                'Input must be set either by an existing input.txt file or by setting via setInput.'
            );
        }

        return $this->input;
    }
}
