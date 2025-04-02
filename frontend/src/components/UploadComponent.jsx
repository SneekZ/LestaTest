import React, { useRef, useState } from "react";

const UploadComponent = ({ setResultData }) => {
  const formRef = useRef(null);
  const fileInputRef = useRef(null);
  const [fileNames, setFileNames] = useState([]);
  const [useLemmatization, setUseLemmatization] = useState(true);
  const [isVisible, setIsVisible] = useState(true);

  const handleFileChange = (e) => {
    const files = Array.from(e.target.files);
    const names = files.map((file) => file.name);
    setFileNames(names);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(formRef.current);
    await getData(formData);
  };

  const handleClear = () => {
    setFileNames([]);
    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }
  };

  const handleLemmatizationToggle = () => {
    setUseLemmatization((prev) => !prev);
  };

  const getData = async (formData) => {
    const endpoint = useLemmatization
      ? "http://localhost:8080/tfidf_lemm"
      : "http://localhost:8080/tfidf";

    const res = await fetch(endpoint, {
      method: "POST",
      body: formData,
    });

    const result = await res.json();
    setResultData(result);
  };

  return (
    <div className="upload-wrapper">
      <button
        className="toggle-upload-button"
        onClick={() => setIsVisible(!isVisible)}
      >
        {isVisible ? "Скрыть загрузку" : "Показать загрузку"}
      </button>

      {isVisible && (
        <form
          className="input-form styled-form"
          ref={formRef}
          onSubmit={handleSubmit}
          encType="multipart/form-data"
        >
          <label htmlFor="file-upload" className="file-label">
            Выбери .txt файлы
            <input
              id="file-upload"
              type="file"
              name="files"
              multiple
              accept=".txt"
              className="file-input"
              ref={fileInputRef}
              onChange={handleFileChange}
            />
          </label>

          <ul className="file-list">
            {fileNames.map((name, index) => (
              <li key={index}>{name}</li>
            ))}
          </ul>

          <div className="lemm-toggle">
            <label className="toggle-label">
              <input
                type="checkbox"
                checked={useLemmatization}
                onChange={handleLemmatizationToggle}
              />
              Лемматизация
            </label>
            <p className="toggle-hint">
              Лемматизация приводит слова к базовой форме, объединяя "модели",
              "моделью", "моделей" в одну — "модель". Это помогает более точно
              оценить важность терминов.
            </p>
          </div>

          <div className="form-buttons">
            <button type="submit" className="submit-button">
              Отправить
            </button>
            <button
              type="button"
              className="clear-button"
              onClick={handleClear}
            >
              Очистить
            </button>
          </div>
        </form>
      )}
    </div>
  );
};

export default UploadComponent;
