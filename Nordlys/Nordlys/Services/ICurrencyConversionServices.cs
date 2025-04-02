using Nordlys.Data;

namespace Nordlys.Services;

/// <summary>
/// Currency conversion services.
/// </summary>
public interface ICurrencyConversionServices
{
    /// <summary>
    /// Converts a given currency to euro.
    /// Returns a negative value if no conversion reference has been registered.
    /// </summary>
    /// <param name="currency">Currency to convert into euro.</param>
    /// <param name="amount">Amount of that currency.</param>
    /// <returns>Amount in euro, or a negative number if no conversion reference has been found.</returns>
    public Task<decimal> ConvertToEuro(Currency currency, decimal amount = 1.0m);
    
}