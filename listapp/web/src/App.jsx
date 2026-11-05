import { useEffect, useState } from "react";

// Thin wrapper around fetch so the component below stays readable.
async function api(path, options = {}) {
  const res = await fetch(path, options);
  if (!res.ok) throw new Error(`${res.status} ${res.statusText}`);
  if (res.status === 204) return null; // DELETE returns no body
  return res.json();
}

const jsonHeaders = { "Content-Type": "application/json" };

export default function App() {
  const [items, setItems] = useState([]);
  const [title, setTitle] = useState("");
  const [error, setError] = useState("");
  const [editingId, setEditingId] = useState(null);
  const [editTitle, setEditTitle] = useState("");

  async function load() {
    try {
      setItems(await api("/api/items"));
      setError("");
    } catch (e) {
      setError(`Cannot reach the API — is the Go server running? (${e.message})`);
    }
  }

  // [] = run once, after the first render.
  useEffect(() => {
    load();
  }, []);

  async function addItem(e) {
    e.preventDefault();
    const clean = title.trim();
    if (!clean) return;
    try {
      await api("/api/items", {
        method: "POST",
        headers: jsonHeaders,
        body: JSON.stringify({ title: clean }),
      });
      setTitle("");
      await load();
    } catch (err) {
      setError(err.message);
    }
  }

  async function toggle(item) {
    try {
      await api(`/api/items/${item.id}`, {
        method: "PUT",
        headers: jsonHeaders,
        body: JSON.stringify({ title: item.title, done: !item.done }),
      });
      await load();
    } catch (err) {
      setError(err.message);
    }
  }

  function startEdit(item) {
    setEditingId(item.id);
    setEditTitle(item.title);
  }

  async function saveEdit(id) {
    const clean = editTitle.trim();
    if (!clean) return;
    const current = items.find((i) => i.id === id);
    try {
      await api(`/api/items/${id}`, {
        method: "PUT",
        headers: jsonHeaders,
        body: JSON.stringify({ title: clean, done: current ? current.done : false }),
      });
      setEditingId(null);
      await load();
    } catch (err) {
      setError(err.message);
    }
  }

  async function remove(id) {
    try {
      await api(`/api/items/${id}`, { method: "DELETE" });
      await load();
    } catch (err) {
      setError(err.message);
    }
  }

  return (
    <main className="app">
      <h1>My List</h1>

      <form onSubmit={addItem} className="add-form">
        <input
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="New item…"
          autoFocus
        />
        <button type="submit">Add</button>
      </form>

      {error && <p className="error">{error}</p>}

      {items.length === 0 && !error && (
        <p className="empty">Nothing here yet — add your first item.</p>
      )}

      <ul className="list">
        {items.map((item) => (
          <li key={item.id} className={item.done ? "item done" : "item"}>
            <input
              type="checkbox"
              checked={item.done}
              onChange={() => toggle(item)}
              aria-label={`Mark ${item.title} as done`}
            />
            {editingId === item.id ? (
              <>
                <input
                  className="edit-input"
                  value={editTitle}
                  onChange={(e) => setEditTitle(e.target.value)}
                  onKeyDown={(e) => {
                    if (e.key === "Enter") saveEdit(item.id);
                  }}
                  autoFocus
                />
                <button onClick={() => saveEdit(item.id)}>Save</button>
                <button className="ghost" onClick={() => setEditingId(null)}>
                  Cancel
                </button>
              </>
            ) : (
              <>
                <span className="title">{item.title}</span>
                <button className="ghost" onClick={() => startEdit(item)}>
                  Edit
                </button>
                <button className="danger" onClick={() => remove(item.id)}>
                  Delete
                </button>
              </>
            )}
          </li>
        ))}
      </ul>
    </main>
  );
}
