<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class UpdateMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        $matakuliahId = $this->route('matakuliah')?->id ?? $this->route('matakuliah');
        return [
            'kode' => ['sometimes', 'string', 'max:10', 'unique:matakuliahs,kode,' . $matakuliahId],
            'nama' => ['sometimes', 'string', 'max:100'],
            'sks' => ['sometimes', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'integer', 'min:1', 'max:8'],
        ];
    }
}