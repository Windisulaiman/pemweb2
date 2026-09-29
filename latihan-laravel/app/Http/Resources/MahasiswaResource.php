<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $allFields = [
            'id' => $this->id,
            'nim' => $this->nim,
            'nama' => $this->nama,
            'email' => $this->email,
            'angkatan' => $this->angkatan,
            'ipk' => (float) $this->ipk,
            'aktif' => $this->aktif,
            'program_studi' => $this->whenLoaded('programStudi', function () {
                return [
                    'id' => $this->programStudi->id,
                    'kode' => $this->programStudi->kode,
                    'nama' => $this->programStudi->nama,
                ];
            }),
            'dibuat_pada' => $this->created_at?->toIso8601String(),
        ];

        // Fitur Dynamic Fields 
        if ($request->has('fields')) {
            $requestedFields = explode(',', $request->query('fields'));
            $filtered = [];

            foreach ($requestedFields as $field) {
                $field = trim($field);
                if (array_key_exists($field, $allFields)) {
                    $filtered[$field] = $allFields[$field];
                }
            }

            return $filtered;
        }

        return $allFields;
    }
}