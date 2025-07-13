interface CheckPayload {
  flightNumber: string;
  date: string;
}

interface GeneratePayload {
  name: string;
  id: string;
  flightNumber: string;
  date: string;
  aircraft: string;
}

export async function check(payload: CheckPayload) {
  const res = await fetch("http://localhost:8080/api/check", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  if (!res.ok) throw new Error("Check failed");
  return res.json();
}

export async function generate(payload: GeneratePayload) {
  const res = await fetch("http://localhost:8080/api/generate", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  const jsonResponse = await res.json();
  if (!res.ok) throw new Error(jsonResponse.message);
  return jsonResponse;
}
