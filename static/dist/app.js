htmx.logger = function(elt, event, data) {
  if(console) {
      console.log(event, elt, data);
  }
}

htmx.onLoad(function(content) {
  // Initialize Leaflet Map
  const map = L.map('map').setView([48.5512778,-123.4139875], 11.2);

  // Add OpenStreetMap Tiles
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
  }).addTo(map);

  // Marker Layer
  const markerLayer = L.layerGroup().addTo(map);

  // Listen to HTMX response and update marker
  document.body.addEventListener('htmx:afterRequest', (event) => {
    if (event.detail.target.id === 'map') {
      const { lat, lng } = JSON.parse(event.detail.xhr.responseText);
      markerLayer.clearLayers(); // Remove existing markers
      L.marker([lat, lng]).addTo(markerLayer);
      map.setView([lat, lng], 8);
    }
  });
})