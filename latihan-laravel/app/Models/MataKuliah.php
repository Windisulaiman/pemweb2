<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class MataKuliah extends Model
{
    protected $table = 'matakuliahs';

    protected $fillable = [
        'kode',
        'nama',
        'sks',
        'semester',
    ];
}