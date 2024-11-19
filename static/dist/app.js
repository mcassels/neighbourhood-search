htmx.logger = function(elt, event, data) {
  if(console) {
      console.log(event, elt, data);
  }
}

htmx.onLoad(function(content) {
  // The map should only load once, on initial page load.
  // All other calls to onLoad will be for swaps and should be ignored here.
  if (!content.querySelector('#map')) {
    return;
  }

  // Initialize Leaflet Map
  const map = L.map('map').setView([48.5512778,-123.4139875], 11.2);

  // Add OpenStreetMap Tiles
  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '© OpenStreetMap contributors'
  }).addTo(map);

  // Marker Layer
  const markerLayer = L.layerGroup().addTo(map);

  // Listen to HTMX response and update marker
  document.body.addEventListener('add-marker', (event) => {
    debugger;
    const { lat, lng } = event.detail;
    L.marker([lat, lng]).addTo(markerLayer);
  });
})