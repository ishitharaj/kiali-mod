import React, { useState } from 'react';

const LLMAssistantWidget: React.FC = () => {
  const [question, setQuestion] = useState('');
  const [response, setResponse] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      const res = await fetch(`/api/llm-assist?q=${encodeURIComponent(question)}`);
      const data = await res.json();
      setResponse(data.answer);
    } catch (error) {
      console.error('Error fetching LLM response:', error);
      setResponse('Error fetching response');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="llm-assistant-widget">
      <button className="chat-icon" onClick={() => setOpen(true)}>Chat</button>
      {open && (
        <div className="chat-drawer">
          <form onSubmit={handleSubmit}>
            <input
              type="text"
              value={question}
              onChange={(e) => setQuestion(e.target.value)}
              placeholder="Ask a question..."
            />
            <button type="submit" disabled={loading}>Submit</button>
          </form>
          {loading ? <div className="spinner">Loading...</div> : <div className="response">{response}</div>}
        </div>
      )}
    </div>
  );
};

export default LLMAssistantWidget; 