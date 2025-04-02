import React from "react";

const TableComponent = ({ resultData }) => {
  if (!resultData || !resultData.tf || !resultData.idf) {
    return <p style={{ color: "#aaa" }}>Нет данных для отображения.</p>;
  }

  const { tf, idf } = resultData;

  const rows = Object.keys(idf).map((word) => {
    const tfValue = tf[word] || 0;
    const idfValue = idf[word];
    const tfidf = tfValue * idfValue;

    return {
      word,
      tf: tfValue,
      idf: idfValue,
      tfidf,
    };
  });

  rows.sort((a, b) => b.idf - a.idf);

  return (
    <div className="result-table">
      <h2>Результаты TF-IDF</h2>
      <div className="table-scroll">
        <table>
          <thead>
            <tr>
              <th>Слово</th>
              <th>TF</th>
              <th>IDF</th>
              <th>TF-IDF</th>
            </tr>
          </thead>
          <tbody>
            {rows.map((row) => (
              <tr key={row.word}>
                <td>{row.word}</td>
                <td>{row.tf.toFixed(5)}</td>
                <td>{row.idf.toFixed(5)}</td>
                <td>{row.tfidf.toFixed(5)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default TableComponent;
