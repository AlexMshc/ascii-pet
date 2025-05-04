import React, {useEffect, useState} from "react";
import {PetProperties, PetService} from "./api";
import "./index.css";

export default function App() {
    const [pet, setPet] = useState<PetProperties | null>(null);
    const [ascii, setAscii] = useState("");
    const [description, setDescription] = useState("");
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        PetService.getPet()
            .then((data) => {
                setPet(data);
            })
            .catch((e) => {
                if (e?.response?.status !== 204) {
                    setError("Не удалось загрузить питомца");
                }
            });
    }, []);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);

        const payload: PetProperties = {ascii, description};
        try {
            await PetService.uploadPet(payload)
            setPet(payload);
        } catch {
            setError("Ошибка при сохранении питомца");
        }
    };

    const handleDelete = async () => {
        setError(null);
        try {
            await PetService.deletePet();
            setPet(null);
            setAscii("");
            setDescription("");
        } catch {
            setError("Не удалось удалить питомца");
        }
    };

    return (
        <div className="container">
            <h1>ASCII Pet App</h1>
            {error && <div className="error">{error}</div>}

            {pet ? (
                <div className="pet-display">
                    <pre className="ascii-art">{pet.ascii}</pre>
                    <p className="description">{pet.description}</p>
                    <button onClick={handleDelete}>Удалить питомца</button>
                </div>
            ) : (
                <div className="no-pet">Питомец не загружен</div>
            )}

            <form onSubmit={handleSubmit} className="upload-form">
                <label>
                    ASCII Art:
                    <textarea
                        value={ascii}
                        onChange={(e) => setAscii(e.target.value)}
                        rows={6}
                        required
                    />
                </label>
                <label>
                    Описание:
                    <input
                        type="text"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        required
                    />
                </label>
                <button type="submit">Сохранить питомца</button>
            </form>
        </div>
    );
}
